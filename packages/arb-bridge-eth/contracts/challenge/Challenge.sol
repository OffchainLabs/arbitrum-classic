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

import "./IChallenge.sol";

import "../vm/IArbBase.sol";

import "../arch/OneStepProof.sol";
import "../arch/Protocol.sol";
import "../arch/Value.sol";

import "../libraries/MerkleLib.sol";


contract Challenge is IChallenge {

    event InitiatedChallenge(
        uint64 deadline
    );

    event ContinuedChallenge (
        uint assertionIndex,
        uint64 deadline
    );

    event BisectedAssertion(
        bytes32[] machineHashes,
        bool[]    didInboxInsns,
        bytes32[] messageAccs,
        bytes32[] logAccs,
        uint64[]  gases,
        uint32 totalSteps,
        uint64 deadline
    );

    event OneStepProofCompleted();

    event TimedOutChallenge (
        bool challengerWrong
    );

    enum State {
        NoChallenge,
        Challenged,
        Bisected
    }

    // Can online initialize once
    string constant CHAL_INIT_STATE = "CHAL_INIT_STATE";
    // Can only continue challenge in response to bisection
    string constant CON_STATE = "CON_STATE";
    // Incorrect previous state
    string constant CON_PREV = "CON_PREV";
    // deadline expired
    string constant CON_DEADLINE = "CON_DEADLINE";
    // Only original challenger can continue challenge
    string constant CON_SENDER = "CON_SENDER";
    // Invalid assertion selected
    string constant CON_PROOF = "CON_PROOF";
    // Can only bisect assertion in response to a challenge
    string constant BIS_STATE = "BIS_STATE";
    // Incorrect previous state
    string constant BIS_INPLEN = "BIS_INPLEN";
    // Incorrect previous state
    string constant BIS_PREV = "BIS_PREV";
    // deadline expired
    string constant BIS_DEADLINE = "BIS_DEADLINE";
    // Only original asserter can continue bisect
    string constant BIS_SENDER = "BIS_SENDER";

    // Can only one step proof following a single step challenge
    string constant OSP_STATE = "OSP_STATE";
    // One step proof with invalid prev state
    string constant OSP_PREV = "OSP_PREV";
    // deadline expired
    string constant OSP_DEADLINE = "OSP_DEADLINE";
    // Proof was incorrect
    string constant OSP_PROOF = "OSP_PROOF";


    address vmAddress;
    // After bisection this is an array of all sub-assertions
    // After a challenge, the first assertion is the main assertion
    bytes32 challengeState;

    uint128[2] escrows;

    address[2] players;

    uint64 deadline;

    // The current deadline at which the challenge timeouts and a winner is
    // declared. This deadline resets at each step in the challenge
    uint32 challengePeriod;

    State state;


    function init(
        address _vmAddress,
        address[2] calldata _players,
        uint128[2] calldata _escrows,
        uint32 _challengePeriod,
        bytes32 _beforeHash,
        bytes32 _beforeInbox,
        uint64[2] calldata _timeBounds,
        bytes32 _assertionHash
    )
        external
    {
        require(state == State.NoChallenge, CHAL_INIT_STATE);

        uint64 newDeadline = uint64(block.number) + uint64(_challengePeriod);
        vmAddress = _vmAddress;
        challengeState = keccak256(
            abi.encodePacked(
                keccak256(
                    abi.encodePacked(
                        _timeBounds[0],
                        _timeBounds[1],
                        _beforeInbox
                    )
                ),
                _beforeHash,
                _assertionHash
            )
        );
        escrows = _escrows;
        players = _players;
        deadline = newDeadline;
        challengePeriod = _challengePeriod;
        state = State.Challenged;

        emit InitiatedChallenge(
            newDeadline
        );
    }

    struct BisectAssertionData {
        bytes32 preData;
        uint64[2] timeBounds;
        bytes32[] machineHashes;
        bool[]    didInboxInsns;
        bytes32[] messageAccs;
        bytes32[] logAccs;
        uint64[] gases;
        uint32 totalSteps;
        uint   bisectionCount;
    }

    function bisectAssertion(
        bytes32 _preData,
        uint64[2] memory _timeBounds,
        bytes32[] memory _machineHashes,
        bool[]    memory _didInboxInsns,
        bytes32[] memory _messageAccs,
        bytes32[] memory _logAccs,
        uint64[] memory _gases,
        uint32 _totalSteps
    )
        public
    {
        _bisectAssertion(
            BisectAssertionData(
                _preData,
                _timeBounds,
                _machineHashes,
                _didInboxInsns,
                _messageAccs,
                _logAccs,
                _gases,
                _totalSteps,
                _machineHashes.length-1
            )
        );
    }

    function _bisectAssertion(BisectAssertionData memory _data) private {
        require(_data.bisectionCount == _data.didInboxInsns.length, BIS_INPLEN);
        require(_data.bisectionCount + 1 == _data.messageAccs.length, BIS_INPLEN);
        require(_data.bisectionCount + 1 == _data.logAccs.length, BIS_INPLEN);
        require(_data.bisectionCount == _data.gases.length, BIS_INPLEN);
        require(State.Challenged == state, BIS_STATE);
        require(block.number <= deadline, BIS_DEADLINE);
        require(msg.sender == players[0], BIS_SENDER);

        uint64 totalGas = 0;
        bool everDidInboxInsn = false;
        for (uint i = 0; i < _data.bisectionCount; i++) {
            totalGas += _data.gases[i];
            everDidInboxInsn = everDidInboxInsn || _data.didInboxInsns[i];
        }

        require(
            keccak256(
                abi.encodePacked(
                    _data.preData,
                    _data.machineHashes[0],
                    Protocol.generateAssertionHash(
                        _data.machineHashes[_data.bisectionCount],
                        everDidInboxInsn,
                        _data.totalSteps,
                        totalGas,
                        _data.messageAccs[0],
                        _data.messageAccs[_data.bisectionCount],
                        _data.logAccs[0],
                        _data.logAccs[_data.bisectionCount]
                    )
                )
            ) == challengeState,
            BIS_PREV
        );

        bytes32[] memory hashes = new bytes32[](_data.bisectionCount);
        hashes[0] = keccak256(
            abi.encodePacked(
                _data.preData,
                _data.machineHashes[0],
                Protocol.generateAssertionHash(
                    _data.machineHashes[1],
                    _data.didInboxInsns[0],
                    _data.totalSteps / uint32(_data.bisectionCount) + _data.totalSteps%uint32(_data.bisectionCount),
                    _data.gases[0],
                    _data.messageAccs[0],
                    _data.messageAccs[1],
                    _data.logAccs[0],
                    _data.logAccs[1]
                )
            )
        );
        for (uint i = 1; i < _data.bisectionCount; i++) {
            if (_data.didInboxInsns[i-1]) {
                _data.preData = keccak256(abi.encodePacked(
                    _data.timeBounds[0],
                    _data.timeBounds[1],
                    Value.hashEmptyTuple()
                ));
            }
            hashes[i] = keccak256(
                abi.encodePacked(
                    _data.preData,
                    _data.machineHashes[i],
                    Protocol.generateAssertionHash(
                        _data.machineHashes[i + 1],
                        _data.didInboxInsns[i],
                        _data.totalSteps / uint32(_data.bisectionCount),
                        _data.gases[i],
                        _data.messageAccs[i],
                        _data.messageAccs[i + 1],
                        _data.logAccs[i],
                        _data.logAccs[i + 1]
                    )
                )
            );
        }
        challengeState = MerkleLib.generateRoot(hashes);
        state = State.Bisected;
        uint64 newDeadline = uint64(block.number) + uint64(challengePeriod);
        deadline = newDeadline;

        emit BisectedAssertion(
            _data.machineHashes,
            _data.didInboxInsns,
            _data.messageAccs,
            _data.logAccs,
            _data.gases,
            _data.totalSteps,
            newDeadline
        );
    }

    function continueChallenge(
        uint _assertionToChallenge,
        bytes memory _proof,
        bytes32 _bisectionRoot,
        bytes32 _bisectionHash
    )
        public
    {
        require(state == State.Bisected, CON_STATE);
        require(_bisectionRoot == challengeState, CON_PREV);
        require(block.number <= deadline, CON_DEADLINE);
        require(msg.sender == players[1], CON_SENDER);
        require(
            MerkleLib.verifyProof(
                _proof,
                _bisectionRoot,
                _bisectionHash,
                _assertionToChallenge + 1
            ),
            CON_PROOF
        );

        state = State.Challenged;
        uint64 newDeadline = uint64(block.number) + uint64(challengePeriod);
        deadline = newDeadline;
        challengeState = _bisectionHash;
        emit ContinuedChallenge(_assertionToChallenge, newDeadline);
    }

    function oneStepProof(
        bytes32 _beforeHash,
        bytes32 _beforeInbox,
        uint64[2] memory _timeBounds,
        bytes32 _afterHash,
        bool    _didInboxInsns,
        bytes32 _firstMessage,
        bytes32 _lastMessage,
        bytes32 _firstLog,
        bytes32 _lastLog,
        uint64  _gas,
        bytes memory _proof
    )
        public
    {
        require(state == State.Challenged, OSP_STATE);
        require(block.number <= deadline, OSP_DEADLINE);
        require(
            keccak256(
                abi.encodePacked(
                    keccak256(
                        abi.encodePacked(
                            _timeBounds[0],
                            _timeBounds[1],
                            _beforeInbox
                        )
                    ),
                    _beforeHash,
                    Protocol.generateAssertionHash(
                        _afterHash,
                        _didInboxInsns,
                        1,
                        _gas,
                        _firstMessage,
                        _lastMessage,
                        _firstLog,
                        _lastLog
                    )
                )
            ) == challengeState,
            OSP_PREV
        );

        uint correctProof = OneStepProof.validateProof(
            _beforeHash,
            _timeBounds,
            _beforeInbox,
            _afterHash,
            _didInboxInsns,
            _firstMessage,
            _lastMessage,
            _firstLog,
            _lastLog,
            _gas,
            _proof
        );

        require(correctProof == 0, OSP_PROOF);
        emit OneStepProofCompleted();
        _asserterWin();
    }

    function asserterTimedOut() public {
        require(
            state == State.Challenged,
            "Can only time out asserter if it is their turn"
        );
        require(block.number > deadline, "Deadline hasn't expired");

        emit TimedOutChallenge(true);
        _challengerWin();
    }

    function challengerTimedOut() public {
        require(
            state == State.Bisected,
            "Can only time out challenger if it is their turn"
        );
        require(block.number > deadline, "Deadline hasn't expired");
        emit TimedOutChallenge(false);
        _asserterWin();
    }

    function _asserterWin() private {
        IArbBase(vmAddress).completeChallenge(
            players,
            [
                escrows[0] + escrows[1] / 2,
                0
            ]
        );
        selfdestruct(msg.sender);
    }

    function _challengerWin() private {
        IArbBase(vmAddress).completeChallenge(
            players,
            [
                0,
                escrows[1] + escrows[0] / 2
            ]
        );
        selfdestruct(msg.sender);
    }
}
