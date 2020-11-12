// SPDX-License-Identifier: Apache-2.0

/*
 * Copyright 2019-2020, Offchain Labs, Inc.
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

import "../challenge/IExecutionChallenge.sol";
import "../challenge/BisectionChallenge.sol";
import "../challenge/ChallengeUtils.sol";

import "../arch/IOneStepProof.sol";

import "../libraries/MerkleLib.sol";

contract ProgressiveExecutionChallenge is IExecutionChallenge, BisectionChallenge {
    using ChallengeUtils for ChallengeUtils.ExecutionAssertion;

    event BisectedAssertion(bytes32[] assertionHashes, uint256 deadlineTicks);

    event OneStepProofCompleted();

    IOneStepProof private executor;

    // Proof was incorrect
    string private constant OSP_PROOF = "OSP_PROOF";

    function connectOneStepProof(address oneStepProof) external {
        executor = IOneStepProof(oneStepProof);
    }

    struct BisectionPrecondition {
        bytes32 beforeMachineHash;
        bytes32 beforeInboxHash;
        bytes32 beforeMessageHash;
        bytes32 beforeLogHash;
    }

    function hash(BisectionPrecondition memory pre) internal pure returns (bytes32) {
        return
            keccak256(
                abi.encodePacked(
                    pre.beforeMachineHash,
                    pre.beforeInboxHash,
                    pre.beforeMessageHash,
                    pre.beforeLogHash
                )
            );
    }

    struct BisectionAssertion {
        uint64 numArbGas;
        bytes32 afterMachineHash;
        bytes32 afterInboxHash;
        bytes32 lastMessageHash;
        uint64 messageCount;
        bytes32 lastLogHash;
        uint64 logCount;
    }

    function hash(BisectionAssertion memory assertion) internal pure returns (bytes32) {
        return
            keccak256(
                abi.encodePacked(
                    assertion.numArbGas,
                    assertion.afterMachineHash,
                    assertion.afterInboxHash,
                    assertion.lastMessageHash,
                    assertion.messageCount,
                    assertion.lastLogHash,
                    assertion.logCount
                )
            );
    }

    function bisectAssertionFirst(
        bytes32 _preconditionHash,
        bytes32 _assertionHash,
        uint64 _numSteps,
        bytes32[] memory _bisectionHashes
    ) public {
        requireMatchesPrevState(
            keccak256(abi.encodePacked(_preconditionHash, _assertionHash, _numSteps))
        );

        _bisectAssertion(_preconditionHash, _assertionHash, _numSteps, _bisectionHashes);
    }

    // _beforeMachine
    //   _machineHash
    //   _inboxAcc
    //   _messageAcc
    //   _logAcc
    function makeBisectionPrecondition(bytes32[4] memory _beforeMachine)
        private
        pure
        returns (BisectionPrecondition memory)
    {
        return
            BisectionPrecondition(
                _beforeMachine[0],
                _beforeMachine[1],
                _beforeMachine[2],
                _beforeMachine[3]
            );
    }

    // _afterMachine
    //   _machineHash
    //   _inboxAcc
    //   _messageAcc
    //   _logAcc

    // _assertionData
    //   _totalSteps
    //   _gasUsed
    //   _messageCount
    //   _logCount
    function makeBisectionAssertion(
        bytes32[4] memory _afterMachine,
        uint64[4] memory _assertionData
    ) private pure returns (BisectionAssertion memory) {
        return
            BisectionAssertion(
                _assertionData[1],
                _afterMachine[0],
                _afterMachine[1],
                _afterMachine[2],
                _assertionData[2],
                _afterMachine[3],
                _assertionData[3]
            );
    }

    function bisectAssertionOther(
        bytes32[4] memory _beforeMachine,
        bytes32[4] memory _afterA1Machine,
        bytes32[4] memory _afterA2Machine,
        uint64[4] memory _assertion1Fields,
        uint64[4] memory _assertion2Fields,
        bytes32[] memory _bisectionHashes
    ) public {
        // steps of A2 >= steps of A1
        require(_assertion2Fields[0] >= _assertion1Fields[0]);
        // gas of A2 >= gas of A1
        require(_assertion2Fields[1] >= _assertion1Fields[1]);
        // message count of A2 >= message count of A1
        require(_assertion2Fields[2] >= _assertion1Fields[2]);
        // log count of A2 >= log count of A1
        require(_assertion2Fields[3] >= _assertion1Fields[3]);

        BisectionPrecondition memory pre = makeBisectionPrecondition(_beforeMachine);

        BisectionAssertion memory a1 = makeBisectionAssertion(_afterA1Machine, _assertion1Fields);
        BisectionAssertion memory a2 = makeBisectionAssertion(_afterA2Machine, _assertion2Fields);

        requireMatchesPrevState(
            keccak256(
                abi.encodePacked(
                    hash(pre),
                    hash(a1),
                    hash(a2),
                    _assertion2Fields[0] - _assertion1Fields[0]
                )
            )
        );

        bytes32 newPreHash = hash(makeBisectionPrecondition(_afterA1Machine));

        uint64[4] memory assertionDiffFields = [
            _assertion2Fields[0] - _assertion1Fields[0],
            _assertion2Fields[1] - _assertion1Fields[1],
            _assertion2Fields[2] - _assertion1Fields[2],
            _assertion2Fields[3] - _assertion1Fields[3]
        ];
        BisectionAssertion memory aDiff = makeBisectionAssertion(
            _afterA2Machine,
            assertionDiffFields
        );

        _bisectAssertion(newPreHash, hash(aDiff), assertionDiffFields[0], _bisectionHashes);
    }

    function _bisectAssertion(
        bytes32 _preHash,
        bytes32 _fullAssertionHash,
        uint64 _totalSteps,
        bytes32[] memory _bisectionHashes
    ) private {
        // require(
        //     bisectionCount == SPLIT_COUNT ||
        //     (_totalSteps < SPLIT_COUNT && bisectionCount == _totalSteps),
        //     "Incorrect bisection count"
        // );

        uint256 bisectionCount = _bisectionHashes.length + 1;
        bytes32[] memory hashes = new bytes32[](bisectionCount);
        hashes[0] = keccak256(
            abi.encodePacked(
                _preHash,
                _bisectionHashes[0],
                uint64(firstSegmentSize(uint256(_totalSteps), bisectionCount))
            )
        );

        uint64 otherStepCount = uint64(otherSegmentSize(uint256(_totalSteps), bisectionCount));
        for (uint256 i = 1; i < bisectionCount - 1; i++) {
            hashes[i] = keccak256(
                abi.encodePacked(
                    _preHash,
                    _bisectionHashes[i - 1],
                    _bisectionHashes[i],
                    otherStepCount
                )
            );
        }
        hashes[bisectionCount] = keccak256(
            abi.encodePacked(
                _preHash,
                _bisectionHashes[bisectionCount - 1],
                _fullAssertionHash,
                otherStepCount
            )
        );

        commitToSegment(hashes);
        asserterResponded();
    }

    // function oneStepProofWithMessage(
    //     bytes32 _firstInbox,
    //     bytes32 _firstMessage,
    //     bytes32 _firstLog,
    //     bytes memory _proof,
    //     uint8 _kind,
    //     uint256 _blockNumber,
    //     uint256 _timestamp,
    //     address _sender,
    //     uint256 _inboxSeqNum,
    //     bytes memory _msgData
    // ) public asserterAction {
    //     (uint64 gas, bytes32[5] memory fields) = executor.executeStepWithMessage(
    //         _firstInbox,
    //         _firstMessage,
    //         _firstLog,
    //         _proof,
    //         _kind,
    //         _blockNumber,
    //         _timestamp,
    //         _sender,
    //         _inboxSeqNum,
    //         _msgData
    //     );

    //     checkProof(gas, _firstInbox, _firstMessage, _firstLog, fields);
    // }

    function oneStepProofFirst(
        bytes32 _firstInbox,
        bytes32 _firstMessage,
        bytes32 _firstLog,
        bytes memory _proof
    ) public asserterAction {
        (uint64 gas, bytes32[5] memory fields) = executor.executeStep(
            _firstInbox,
            _firstMessage,
            _firstLog,
            _proof
        );

        (bytes32 preconditionHash, bytes32 assertionHash) = calculateProof(
            gas,
            _firstInbox,
            _firstMessage,
            _firstLog,
            fields
        );

        requireMatchesPrevState(
            keccak256(abi.encodePacked(preconditionHash, assertionHash, uint64(1)))
        );

        emit OneStepProofCompleted();
        _asserterWin();
    }

    // function oneStepProofOther(
    //     bytes32 _firstInbox,
    //     bytes32 _firstMessage,
    //     bytes32 _firstLog,
    //     bytes memory _proof
    // ) public asserterAction {
    //     (uint64 gas, bytes32[5] memory fields) = executor.executeStep(
    //         _firstInbox,
    //         _firstMessage,
    //         _firstLog,
    //         _proof
    //     );

    //     (bytes32 preconditionHash, bytes32 assertionsHash) = calculateProof(gas, _firstInbox, _firstMessage, _firstLog, fields);

    //     requireMatchesPrevState(
    //         keccak256(abi.encodePacked(preconditionHash, assertionHash, 1))
    //     );

    //     emit OneStepProofCompleted();
    //     _asserterWin();
    // }

    // fields
    //  startMachineHash
    //  endMachineHash
    //  afterInboxHash
    //  afterMessagesHash
    //  afterLogsHash

    function calculateProof(
        uint64 gas,
        bytes32 firstInbox,
        bytes32 firstMessage,
        bytes32 firstLog,
        bytes32[5] memory fields
    ) private pure returns (bytes32, bytes32) {
        bytes32 preconditionHash = hash(
            BisectionPrecondition(fields[0], firstInbox, firstMessage, firstLog)
        );

        // The one step proof already guarantees us that firstMessage and lastMessage
        // are either one or 0 messages apart and the same is true for logs. Therefore
        // we can infer the message count and log count based on whether the fields
        // are equal or not
        bytes32 assertionHash = hash(
            BisectionAssertion(
                gas,
                fields[1],
                fields[2],
                fields[3],
                firstMessage == fields[3] ? 0 : 1,
                fields[4],
                firstLog == fields[4] ? 0 : 1
            )
        );
        return (preconditionHash, assertionHash);
    }
}
