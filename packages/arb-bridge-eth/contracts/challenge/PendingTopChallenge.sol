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

import "./BisectionChallenge.sol";
import "./IPendingTopChallenge.sol";

import "../arch/Protocol.sol";

contract PendingTopChallenge is BisectionChallenge, IPendingTopChallenge {

    event Bisected(
        bytes32[] chainHashes,
        uint32 totalLength,
        uint64 deadline
    );

    event OneStepProofCompleted();

    // Proof was incorrect
    string constant HC_OSP_PROOF = "HC_OSP_PROOF";

    function init(
        address _vmAddress,
        address payable _asserter,
        address payable _challenger,
        uint32 _challengePeriod,
        bytes32 _topHash,
        bytes32 _lowerHash
    )
        external
    {
        BisectionChallenge.initializeBisection(
            _vmAddress,
            _asserter,
            _challenger,
            _challengePeriod,
            encodeSegment(_topHash, _lowerHash, 0)
        );
    }

    function bisectFirst(
        bytes32[] memory _chainHashes,
        uint32 _chainLength
    )
        public
        asserterAction
    {
        uint bisectionCount = _chainHashes.length - 1;

        requireMatchesPrevState(
            encodeSegment(
                _chainHashes[0],
                _chainHashes[bisectionCount],
                0
            )
        );

        _bisect(_chainHashes, _chainLength);
    }

    function bisect(
        bytes32[] memory _chainHashes,
        uint32 _chainLength
    )
        public
        asserterAction
    {
        uint bisectionCount = _chainHashes.length - 1;

        requireMatchesPrevState(
            encodeSegment(
                _chainHashes[0],
                _chainHashes[bisectionCount],
                _chainLength
            )
        );

        _bisect(_chainHashes, _chainLength);
    }

    function _bisect(
        bytes32[] memory _chainHashes,
        uint32 _chainLength
    )
        internal
    {
        require(_chainLength > 1, "Can't bisect chain of less than 2");
        uint bisectionCount = _chainHashes.length - 1;
        bytes32[] memory hashes = new bytes32[](bisectionCount);
        hashes[0] = encodeSegment(
            _chainHashes[0],
            _chainHashes[1],
            firstSegmentSize(_chainLength, bisectionCount)
        );
        for (uint i = 1; i < bisectionCount; i++) {
            hashes[i] = encodeSegment(
                _chainHashes[i],
                _chainHashes[i + 1],
                otherSegmentSize(_chainLength, bisectionCount)
            );
        }

        commitToSegment(hashes);
        asserterResponded();
        emit Bisected(
            _chainHashes,
            _chainLength,
            deadline
        );
    }

    function oneStepProof(
        bytes32 _topHash,
        bytes32 _lowerHash,
        bytes32 _value
    )
        public
        asserterAction
    {
        requireMatchesPrevState(
            encodeSegment(
                _topHash,
                _lowerHash,
                uint32(1)
            )
        );

        require(Protocol.addMessageToPending(_lowerHash, _value) == _topHash, HC_OSP_PROOF);

        emit OneStepProofCompleted();
        _asserterWin();
    }

    function encodeSegment(
        bytes32 _topHash,
        bytes32 _lowerHash,
        uint32 _chainLength
    )
        private
        pure
        returns(bytes32)
    {
        return keccak256(
            abi.encodePacked(
                _topHash,
                _lowerHash,
                _chainLength
            )
        );
    }
}
